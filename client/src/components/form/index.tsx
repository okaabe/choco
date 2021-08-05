import React from 'react';

import {
    FormProperties
} from './form.props';

import {
    FormContainer,
    FormInputContainer,
    Input
} from './form.styles';

export const Form: React.FC<FormProperties> = ({
    inputs,
    submitButton,
    children,
}) => {
    return (
        <FormContainer>
            <FormInputContainer>
                {inputs.map(({
                    stateDispatcher,
                    isPassword,
                    ...inputProperties
                }) => {
                    return isPassword ? <Input 
                        type="password"
                        onChange={(e) => stateDispatcher!(e.target.value)}
                        {...inputProperties}
                    /> : <Input 
                        onChange={ (e) => stateDispatcher!(e.target.value) }
                        {...inputProperties}
                    />
                })}
            </FormInputContainer>
        </FormContainer>
    )
}