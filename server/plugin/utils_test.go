package plugin

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/mattermost/mattermost-plugin-azure-devops/mocks"
	"github.com/mattermost/mattermost-plugin-azure-devops/server/config"
	"github.com/mattermost/mattermost-plugin-azure-devops/server/constants"
	"github.com/mattermost/mattermost-plugin-azure-devops/server/serializers"
	"github.com/mattermost/mattermost-plugin-azure-devops/server/testutils"
)

type mockBLock struct{}

func (b *mockBLock) BlockSize() int { return 1 }

func (b *mockBLock) Encrypt(_, _ []byte) {}

func (b *mockBLock) Decrypt(_, _ []byte) {}

type mockAesgcm struct{}

func (a *mockAesgcm) NonceSize() int { return 1 }

func (a *mockAesgcm) Overhead() int { return 1 }

func (a *mockAesgcm) Seal(dst, nonce, plaintext, additionalData []byte) []byte { return []byte("mock") }

func (a *mockAesgcm) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	return []byte("mock"), nil
}

func TestSendEphemeralPostForCommand(t *testing.T) {
	p := Plugin{}
	mockAPI := &plugintest.API{}
	p.API = mockAPI
	for _, testCase := range []struct {
		description string
		text        string
		args        model.CommandArgs
	}{
		{
			description: "SendEphemeralPostForCommand: valid",
			text:        "mockText",
			args: model.CommandArgs{
				UserId:    testutils.MockMattermostUserID,
				ChannelId: testutils.MockChannelID,
			},
		},
	} {
		mockAPI.On("SendEphemeralPost", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Post")).Return(nil)

		t.Run(testCase.description, func(t *testing.T) {
			resp, _ := p.sendEphemeralPostForCommand(&testCase.args, testCase.text)
			assert.NotNil(t, resp)
		})
	}
}

func TestDM(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description      string
		mattermostUserID string
		format           string
		args             model.CommandArgs
		post             *model.Post
		postErr          *model.AppError
		channel          *model.Channel
		channelErr       *model.AppError
	}{
		{
			description:      "DM: valid",
			mattermostUserID: testutils.MockMattermostUserID,
			format:           "mockFormat",
			args: model.CommandArgs{
				UserId:    testutils.MockMattermostUserID,
				ChannelId: testutils.MockChannelID,
			},
			channel: &model.Channel{
				Id:   testutils.MockChannelID,
				Type: model.CHANNEL_OPEN,
			},
			channelErr: nil,
			post: &model.Post{
				Id: "mockID",
			},
			postErr: nil,
		},
		{
			description:      "DM: with channelErr",
			mattermostUserID: testutils.MockMattermostUserID,
			format:           "mockFormat",
			args: model.CommandArgs{
				UserId:    testutils.MockMattermostUserID,
				ChannelId: testutils.MockChannelID,
			},
			channelErr: &model.AppError{
				Message:       "failed to get direct channel",
				DetailedError: "failed to get direct channel",
			},
		},
		{
			description:      "DM: with postErr",
			mattermostUserID: testutils.MockMattermostUserID,
			format:           "mockFormat",
			args: model.CommandArgs{
				UserId:    testutils.MockMattermostUserID,
				ChannelId: testutils.MockChannelID,
			},
			channel: &model.Channel{
				Id:   testutils.MockChannelID,
				Type: model.CHANNEL_OPEN,
			},
			channelErr: nil,
			postErr: &model.AppError{
				Message:       "failed to create post",
				DetailedError: "failed to create post",
			},
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			mockAPI := &plugintest.API{}
			p.API = mockAPI

			mockAPI.On("LogError", testutils.GetMockArgumentsWithType("string", 5)...)
			mockAPI.On("GetDirectChannel", testutils.GetMockArgumentsWithType("string", 2)...).Return(testCase.channel, testCase.channelErr)
			mockAPI.On("CreatePost", mock.AnythingOfType("*model.Post")).Return(testCase.post, testCase.postErr)

			resp, _ := p.DM(testCase.mattermostUserID, testCase.format, false, &testCase.args)
			assert.NotNil(t, resp)
		})
	}
}

func TestEncode(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
		encrypted   string
	}{
		{
			description: "Encode: valid",
			encrypted:   "mockData",
		},
		{
			description: "Encode: empty encrypted string",
			encrypted:   "",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp := p.Encode([]byte(testCase.encrypted))
			assert.NotNil(t, resp)
		})
	}
}

func TestDecode(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description   string
		encoded       string
		decodingError error
	}{
		{
			description: "Decode: valid",
			encoded:     "mockData",
		},
		{
			description: "Decode: empty encoded string",
			encoded:     "",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp, err := p.Decode(testCase.encoded)
			assert.NotNil(t, resp)
			assert.Nil(t, err)
		})
	}
}

func TestEncrypt(t *testing.T) {
	defer monkey.UnpatchAll()
	p := Plugin{}
	for _, testCase := range []struct {
		description    string
		expectedError  string
		newCipherError error
		newGCMError    error
		readFullError  error
		plain          string
		secret         string
	}{
		{
			description: "Encrypt: length of secret equal to 0",
			plain:       "mockPlain",
		},
		{
			description:    "Encrypt: aes.NewCipher give error",
			secret:         "mockSecret",
			expectedError:  "newCipherError",
			newCipherError: errors.New("newCipherError"),
		},
		{
			description:   "Encrypt: cipher.NewGCM give error",
			secret:        "mockSecret",
			expectedError: "newGCMError",
			newGCMError:   errors.New("newGCMError"),
		},
		{
			description:   "Encrypt: io.ReadFull give error",
			secret:        "mockSecret",
			expectedError: "readFullError",
			readFullError: errors.New("readFullError"),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			monkey.Patch(aes.NewCipher, func(a []byte) (cipher.Block, error) {
				return &mockBLock{}, testCase.newCipherError
			})
			monkey.Patch(cipher.NewGCM, func(_ cipher.Block) (cipher.AEAD, error) {
				return &mockAesgcm{}, testCase.newGCMError
			})
			monkey.Patch(io.ReadFull, func(_ io.Reader, _ []byte) (int, error) {
				return 1, testCase.readFullError
			})
			resp, err := p.Encrypt([]byte(testCase.plain), []byte(testCase.secret))
			if testCase.expectedError != "" {
				assert.Nil(t, resp)
				assert.NotNil(t, err)
				return
			}

			assert.NotNil(t, resp)
			assert.Nil(t, err)
		})
	}
}

func TestDecrypt(t *testing.T) {
	defer monkey.UnpatchAll()
	p := Plugin{}
	for _, testCase := range []struct {
		description    string
		expectedError  string
		newCipherError error
		newGCMError    error
		resdFullError  error
		encrypted      string
		secret         string
	}{
		{
			description: "Decrypt: length of secret equal to 0",
			encrypted:   "mockPlain",
		},
		{
			description:    "Decrypt: aes.NewCipher give error",
			secret:         "mockSecret",
			expectedError:  "newCipherError",
			newCipherError: errors.New("newCipherError"),
		},
		{
			description:   "Decrypt: cipher.NewGCM give error",
			secret:        "mockSecret",
			expectedError: "newGCMError",
			newGCMError:   errors.New("newGCMError"),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			monkey.Patch(aes.NewCipher, func([]byte) (cipher.Block, error) {
				return &mockBLock{}, testCase.newCipherError
			})
			monkey.Patch(cipher.NewGCM, func(cipher.Block) (cipher.AEAD, error) {
				return &mockAesgcm{}, testCase.newGCMError
			})
			resp, err := p.Decrypt([]byte(testCase.encrypted), []byte(testCase.secret))
			if testCase.expectedError != "" {
				assert.Nil(t, resp)
				assert.NotNil(t, err)
				return
			}

			assert.NotNil(t, resp)
			assert.Nil(t, err)
		})
	}
}

func TestGetSiteURL(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
	}{
		{
			description: "GetSiteURL: valid",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			p.setConfiguration(
				&config.Configuration{
					MattermostSiteURL: "mockMattermostSiteURL",
				})
			resp := p.GetSiteURL()
			assert.NotNil(t, resp)
		})
	}
}

func TestGetPluginURLPath(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
	}{
		{
			description: "GetPluginURLPath: valid",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp := p.GetPluginURLPath()
			assert.NotNil(t, resp)
		})
	}
}

func TestGetPluginURL(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
	}{
		{
			description: "GetPluginURLPath: valid",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp := p.GetPluginURL()
			assert.NotNil(t, resp)
		})
	}
}

func TestSanitizeURLPath(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description           string
		organization          string
		project               string
		otherPathInput        string
		isAnyPathInputInvalid bool
	}{
		{
			description:    "SanitizeURLPaths: valid organization, project and otherInputs",
			organization:   "dummyOrg",
			project:        "dummy_project",
			otherPathInput: "dummy_path",
		},
		{
			description:           "SanitizeURLPaths: invalid organization",
			organization:          "dummy_org",
			project:               "",
			otherPathInput:        "",
			isAnyPathInputInvalid: true,
		},
		{
			description:           "SanitizeURLPaths: invalid project",
			organization:          "",
			project:               "../dummy_project",
			otherPathInput:        "",
			isAnyPathInputInvalid: true,
		},
		{
			description:           "SanitizeURLPaths: invalid project with escaped chars",
			organization:          "",
			project:               "%5c..%5c..dummy_project",
			otherPathInput:        "",
			isAnyPathInputInvalid: true,
		},
		{
			description:           "SanitizeURLPaths: invalid otherInputs",
			organization:          "",
			project:               "",
			otherPathInput:        "dummy_path/../",
			isAnyPathInputInvalid: true,
		},
		{
			description:           "SanitizeURLPaths: invalid otherInputs with escaped chars",
			organization:          "",
			project:               "",
			otherPathInput:        "dummy_path%2f%2e.",
			isAnyPathInputInvalid: true,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			_, err := p.SanitizeURLPaths(testCase.organization, testCase.project, testCase.otherPathInput)
			if testCase.isAnyPathInputInvalid {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseAuthToken(t *testing.T) {
	defer monkey.UnpatchAll()
	p := Plugin{}
	for _, testCase := range []struct {
		description    string
		expectedError  string
		decodedToken   []byte
		decodeError    error
		decryptedToken []byte
		decryptError   error
		encodedToken   string
	}{
		{
			description:    "ParseAuthToken: token is parsed successfully",
			encodedToken:   "mockEncodedToken",
			decodedToken:   []byte("mockDecodedToken"),
			decryptedToken: []byte("mockDecryptedToken"),
		},
		{
			description:   "ParseAuthToken: token is not decoded successfully",
			expectedError: "error decoding oAuth token",
			decodeError:   errors.New("error decoding oAuth token"),
			encodedToken:  "mockEncodedToken",
		},
		{
			description:   "ParseAuthToken: token is not decrypted successfully",
			expectedError: "error decrypting oAuth token",
			decodedToken:  []byte("mockDecryptedToken"),
			decryptError:  errors.New("error decrypting oAuth token"),
			encodedToken:  "mockEncodedToken",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "Decode", func(_ *Plugin, _ string) ([]byte, error) {
				return testCase.decodedToken, testCase.decodeError
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "Decrypt", func(_ *Plugin, _, _ []byte) ([]byte, error) {
				return testCase.decryptedToken, testCase.decryptError
			})

			res, err := p.ParseAuthToken(testCase.encodedToken)

			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
				assert.Equal(t, "", res)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, res)
		})
	}
}

func TestAddAuthorization(t *testing.T) {
	p := Plugin{}
	mockAPI := &plugintest.API{}
	mockCtrl := gomock.NewController(t)
	mockedStore := mocks.NewMockKVStore(mockCtrl)
	p.API = mockAPI
	p.Store = mockedStore
	for _, testCase := range []struct {
		description       string
		user              *serializers.User
		token             string
		parseAuthTokenErr error
		loadUserErr       error
	}{
		{
			description: "AddAuthorization: valid",
			user: &serializers.User{
				AccessToken: "mockAccessToken",
				UserProfile: serializers.UserProfile{
					ID: testutils.MockAzureDevopsUserID,
				},
			},
			token: "mockToken",
		},
		{
			description: "AddAuthorization: error while loading user",
			loadUserErr: errors.New("mockError"),
		},
		{
			description: "AddAuthorization: empty user",
			user: &serializers.User{
				AccessToken: "mockAccessToken",
				UserProfile: serializers.UserProfile{
					ID: testutils.MockAzureDevopsUserID,
				},
			},
			parseAuthTokenErr: errors.New("mockError"),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			mockedStore.EXPECT().LoadAzureDevopsUserIDFromMattermostUser(testutils.MockMattermostUserID).Return(testutils.MockAzureDevopsUserID, nil)

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "ParseAuthToken", func(_ *Plugin, _ string) (string, error) {
				return testCase.token, testCase.parseAuthTokenErr
			})

			mockedStore.EXPECT().LoadAzureDevopsUserDetails(testutils.MockAzureDevopsUserID).Return(testCase.user, testCase.loadUserErr)

			req := httptest.NewRequest(http.MethodGet, "/mockURL", bytes.NewBufferString(`{}`))
			resp := p.AddAuthorization(req, testutils.MockMattermostUserID)
			if testCase.loadUserErr != nil || testCase.parseAuthTokenErr != nil {
				assert.NotNil(t, resp)
				return
			}

			assert.Nil(t, resp)
		})
	}
}

func TestIsProjectLinked(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
		projectList []serializers.ProjectDetails
		project     serializers.ProjectDetails
	}{
		{
			description: "IsProjectLinked: project present in project list",
			projectList: []serializers.ProjectDetails{
				{
					ProjectName:      testutils.MockProjectName,
					OrganizationName: testutils.MockOrganization,
				},
			},
			project: serializers.ProjectDetails{
				ProjectName:      testutils.MockProjectName,
				OrganizationName: testutils.MockOrganization,
			},
		},
		{
			description: "IsProjectLinked: project not present in project list",
			projectList: []serializers.ProjectDetails{
				{
					ProjectName:      testutils.MockProjectName,
					OrganizationName: testutils.MockOrganization,
				},
			},
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp, isProjectLinked := p.IsProjectLinked(testCase.projectList, testCase.project)
			if isProjectLinked {
				assert.NotNil(t, resp)
				return
			}

			assert.Nil(t, resp)
		})
	}
}

func TestIsSubscriptionPresent(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description      string
		subscriptionList []*serializers.SubscriptionDetails
		subscription     *serializers.SubscriptionDetails
	}{
		{
			description:      "test IsSubscriptionPresent with subscription present in subscription list",
			subscriptionList: testutils.GetSuscriptionDetailsPayload(testutils.MockMattermostUserID, testutils.MockServiceType, testutils.MockEventType),
			subscription:     testutils.GetSuscriptionDetailsPayload(testutils.MockMattermostUserID, testutils.MockServiceType, testutils.MockEventType)[0],
		},
		{
			description:      "test IsSubscriptionPresent with subscription not present in subscription list",
			subscriptionList: testutils.GetSuscriptionDetailsPayload(testutils.MockMattermostUserID, testutils.MockServiceType, testutils.MockEventType),
			subscription:     &serializers.SubscriptionDetails{},
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp, isSubscriptionPresent := p.IsSubscriptionPresent(testCase.subscriptionList, testCase.subscription)
			if isSubscriptionPresent {
				assert.NotNil(t, resp)
				return
			}

			assert.Nil(t, resp)
		})
	}
}

func TestIsAnyProjectLinked(t *testing.T) {
	p := Plugin{}
	mockAPI := &plugintest.API{}
	p.API = mockAPI
	for _, testCase := range []struct {
		description      string
		mattermostUserID string
		projectList      []serializers.ProjectDetails
		projectErr       error
	}{
		{
			description:      "IsAnyProjectLinked: valid",
			mattermostUserID: testutils.MockMattermostUserID,
			projectList: []serializers.ProjectDetails{
				{
					ProjectName:      testutils.MockProjectName,
					OrganizationName: testutils.MockOrganization,
				},
			},
		},
		{
			description:      "IsAnyProjectLinked: empty project list",
			mattermostUserID: testutils.MockMattermostUserID,
		},
		{
			description:      "IsAnyProjectLinked: error while getting project list",
			mattermostUserID: testutils.MockMattermostUserID,
			projectErr:       errors.New("mockError"),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockedStore := mocks.NewMockKVStore(mockCtrl)
			mockedStore.EXPECT().GetAllProjects(testCase.mattermostUserID).Return(testCase.projectList, testCase.projectErr)
			p.Store = mockedStore
			isAnyProjectLinked, err := p.IsAnyProjectLinked(testCase.mattermostUserID)
			if testCase.projectErr != nil {
				assert.Empty(t, isAnyProjectLinked)
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			if testCase.projectList != nil {
				assert.Equal(t, true, isAnyProjectLinked)
				return
			}

			assert.Equal(t, false, isAnyProjectLinked)
		})
	}
}

func TestGetConnectAccountFirstMessage(t *testing.T) {
	p := Plugin{}
	for _, testCase := range []struct {
		description string
	}{
		{
			description: "GetConnectAccountFirstMessage: valid",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			resp := p.getConnectAccountFirstMessage()
			assert.NotNil(t, resp)
		})
	}
}

func TestGetOffsetAndLimitFromQueryParams(t *testing.T) {
	p := Plugin{}
	mockAPI := &plugintest.API{}
	p.API = mockAPI
	for _, testCase := range []struct {
		description       string
		queryParamPage    string
		queryParamPerPage string
		expectedOffset    int
		expectedLimit     int
	}{
		{
			description:       "GetOffsetAndLimitFromQueryParams: valid page and per_page query params",
			queryParamPage:    "1",
			queryParamPerPage: "10",
			expectedOffset:    10,
			expectedLimit:     10,
		},
		{
			description:       "GetOffsetAndLimitFromQueryParams: empty page and per_page query params",
			queryParamPage:    "",
			queryParamPerPage: "",
			expectedOffset:    0,
			expectedLimit:     constants.DefaultPerPageLimit,
		},
		{
			description:       "GetOffsetAndLimitFromQueryParams: invalid page query param",
			queryParamPage:    "invalidNonIntegerString",
			queryParamPerPage: "10",
			expectedOffset:    0,
			expectedLimit:     10,
		},
		{
			description:       "GetOffsetAndLimitFromQueryParams: invalid per_page query param",
			queryParamPage:    "1",
			queryParamPerPage: "invalidNonIntegerString",
			expectedOffset:    constants.DefaultPerPageLimit,
			expectedLimit:     constants.DefaultPerPageLimit,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			if testCase.queryParamPage != "1" && testCase.queryParamPerPage != "10" && testCase.expectedLimit != 10 && testCase.expectedOffset != 10 {
				mockAPI.On("LogError", testutils.GetMockArgumentsWithType("string", 3)...)
				defer mockAPI.AssertExpectations(t)
			}

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/mockTeamID?project=%s&page=%s&per_page=%s", constants.PathGetSubscriptions, testutils.MockProjectName, testCase.queryParamPage, testCase.queryParamPerPage), bytes.NewBufferString(`{}`))
			req.Header.Add(constants.HeaderMattermostUserID, testutils.MockMattermostUserID)

			offset, limit := p.GetOffsetAndLimitFromQueryParams(req)

			assert.Equal(t, testCase.expectedOffset, offset)
			assert.Equal(t, testCase.expectedLimit, limit)
		})
	}
}

func TestParseSubscriptionsToCommandResponse(t *testing.T) {
	defer monkey.UnpatchAll()
	p := Plugin{}
	mockAPI := &plugintest.API{}
	p.API = mockAPI
	for _, testCase := range []struct {
		description       string
		subscriptionsList []*serializers.SubscriptionDetails
		command           string
		expectedMessage   string
		createdBy         string
		err               error
	}{
		{
			description:       "ParseSubscriptionsToCommandResponse: empty repos subscription list",
			command:           constants.CommandRepos,
			subscriptionsList: []*serializers.SubscriptionDetails{},
			expectedMessage:   fmt.Sprintf("No %s subscription exists", constants.CommandRepos),
		},
		{
			description:       "ParseSubscriptionsToCommandResponse: empty boards subscription list",
			command:           constants.CommandBoards,
			subscriptionsList: []*serializers.SubscriptionDetails{},
			expectedMessage:   fmt.Sprintf("No %s subscription exists", constants.CommandBoards),
		},
		{
			description:       "ParseSubscriptionsToCommandResponse: error in fetching filtered subscription list",
			command:           constants.CommandBoards,
			subscriptionsList: []*serializers.SubscriptionDetails{},
			expectedMessage:   constants.GenericErrorMessage,
			err:               errors.New("error in fetching filtered subscription list"),
		},
		{
			description:       "ParseSubscriptionsToCommandResponse: subscriptions created by the user",
			command:           constants.CommandBoards,
			subscriptionsList: testutils.GetSuscriptionDetailsPayload(testutils.MockMattermostUserID, constants.CommandBoards, constants.SubscriptionEventWorkItemCreated),
			createdBy:         constants.FilterCreatedByMe,
			expectedMessage:   fmt.Sprintf("###### %s subscription(s)\n| Subscription ID | Organization | Project | Event Type | Created By | Channel |\n| :-------------- | :----------- | :------ | :--------- | :--------- | :------ |\n| mockSubscriptionID | mockOrganization | mockProjectName | Work Item Created | mockCreatedBy | mockChannelName |\n", cases.Title(language.Und).String(constants.CommandBoards)),
		},
		{
			description:       "ParseSubscriptionsToCommandResponse: subscriptions created by anyone",
			command:           constants.CommandBoards,
			subscriptionsList: testutils.GetSuscriptionDetailsPayload(testutils.MockMattermostUserID, constants.CommandBoards, constants.SubscriptionEventWorkItemCreated),

			createdBy:       constants.FilterCreatedByAnyone,
			expectedMessage: fmt.Sprintf("###### %s subscription(s)\n| Subscription ID | Organization | Project | Event Type | Created By | Channel |\n| :-------------- | :----------- | :------ | :--------- | :--------- | :------ |\n| mockSubscriptionID | mockOrganization | mockProjectName | Work Item Created | mockCreatedBy | mockChannelName |\n", cases.Title(language.Und).String(constants.CommandBoards)),
		},
		{
			description:       "ParseSubscriptionsToCommandResponse: no subscriptions created by the user is present",
			command:           constants.CommandBoards,
			subscriptionsList: testutils.GetSuscriptionDetailsPayload("mockUserID-2", constants.CommandBoards, constants.SubscriptionEventWorkItemCreated),
			createdBy:         constants.FilterCreatedByMe,
			expectedMessage:   fmt.Sprintf("No %s subscription exists", constants.CommandBoards),
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			mockAPI.On("LogError", testutils.GetMockArgumentsWithType("string", 3)...)

			monkey.PatchInstanceMethod(reflect.TypeOf(&p), "GetSubscriptionsForAccessibleChannelsOrProjects", func(_ *Plugin, _ []*serializers.SubscriptionDetails, _, _, _ string) ([]*serializers.SubscriptionDetails, error) {
				return testCase.subscriptionsList, testCase.err
			})

			message := p.ParseSubscriptionsToCommandResponse(testCase.subscriptionsList, testutils.MockChannelID, testCase.createdBy, testutils.MockMattermostUserID, testCase.command, "mockTeamID")
			assert.Equal(t, testCase.expectedMessage, message)
		})
	}
}
