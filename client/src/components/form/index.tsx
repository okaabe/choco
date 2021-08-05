import React from 'react';

import { FormInput, FormProperties } from './form.types';
import { FormContainer, Input } from './form.styles';

export const Form: React.FC<FormProperties> = ({
    inputs
}) => {
    return (
        <FormContainer>
            { inputs?.map(({ onChangeState, ...props }) => 
                <Input
                    onChange={(e) => onChangeState!(e.target.value)}
                    { ...props }
                />
            )}
        </FormContainer>
    )
}