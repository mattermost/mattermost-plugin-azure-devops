import React from 'react';

import Dropdown from 'components/dropdown';
import Input from 'components/inputField';

type Props = {
    fieldConfig: Pick<ModalFormFieldConfig, 'label' | 'type' | 'validations' | 'isFieldDisabled'>
    value: string | null
    optionsList?: LabelValuePair[]
    onChange: (newValue: string, targetLabel?: string, selectedOption?: Record<string, string>) => void;
    error?: string
    isDisabled?: boolean
}

/**
 * A generic component to render form
 * you can add multiple input field types here
 */
const Form = ({fieldConfig: {label, type, validations, isFieldDisabled}, value, optionsList, onChange, error, isDisabled}: Props): JSX.Element => {
    switch (type) {
    case 'dropdown' :
        return (
            <Dropdown
                placeholder={label}
                value={value}
                onChange={onChange}
                options={optionsList || []}
                required={validations?.isRequired as boolean}
                error={error}
                disabled={isDisabled || isFieldDisabled}
            />
        );
    case 'text' :
        return (
            <Input
                type='text'
                placeholder={label}
                value={value ?? ''}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => onChange(e.target.value)}
                error={error}
                required={validations?.isRequired as boolean}
                disabled={isDisabled || isFieldDisabled}
            />
        );
    default:
        return <></>;
    }
};

export default Form;
