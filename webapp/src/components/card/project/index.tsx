import React from 'react';

import BaseCard from 'components/card/base';
import IconButton from 'components/buttons/iconButton';

import {onPressingEnterKey} from 'utils';
import SVGWrapper from 'components/svgWrapper';
import plugin_constants from 'plugin_constants';
import LabelValuePair from 'components/labelValuePair';

type ProjectCardProps = {
    onProjectTitleClick: (projectDetails: ProjectDetails) => void
    handleUnlinkProject: (projectDetails: ProjectDetails) => void
    projectDetails: ProjectDetails
}

const ProjectCard = ({onProjectTitleClick, projectDetails: {organizationName, projectName}, projectDetails, handleUnlinkProject}: ProjectCardProps) => (
    <BaseCard>
        <div className='d-flex'>
            <div className='project-details'>
                <LabelValuePair
                    label={
                        <SVGWrapper
                            width={16}
                            height={16}
                            viewBox='0 0 14 12'
                        >
                            {plugin_constants.SVGIcons.project}
                        </SVGWrapper>
                    }
                    onClickValue={() => onProjectTitleClick(projectDetails)}
                    value={projectName}
                    labelExtraClassName='margin-top-1'
                />
                <LabelValuePair
                    label={
                        <SVGWrapper
                            width={13}
                            height={13}
                            viewBox='0 0 10 10'
                        >
                            {plugin_constants.SVGIcons.organization}
                        </SVGWrapper>
                    }
                    value={organizationName}
                />
            </div>
            <div className='button-wrapper'>
                <IconButton
                    tooltipText='Unlink project'
                    iconClassName='fa fa-chain-broken'
                    extraClass='unlink-button'
                    onClick={() => handleUnlinkProject(projectDetails)}
                />
            </div>
        </div>
    </BaseCard>
);

export default ProjectCard;

