import React from 'react';

import { View } from '../../../components/view';
import { SigninContent } from './signin.styles';
import { Form } from '../../../components/form';

export const SignIn: React.FC = () => {
    return (
        <View>
            <SigninContent>
                <Form 
                    inputs={[
                        {
                            id: "username",
                            name: "username",
                            placeholder: "turururu"
                        }
                    ]}
                    submitButton={{
                        message: "Sign In",
                        onPress: () => {}
                    }}
                />
            </SigninContent>
        </View>
    )
}