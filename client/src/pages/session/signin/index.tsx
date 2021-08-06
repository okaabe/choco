import React, { useState } from 'react';

import { View } from '../../../components/view';
import { Form } from '../../../components/form';
import { SigninContent } from './signin.styles';

export const SignIn: React.FC = () => {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    return (
        <View>
            <SigninContent>
                <Form inputs={[
                    {
                        placeholder: "Username",
                        onChangeState: setUsername,
                    },
                    {
                        type: "password",
                        placeholder: "Password",
                        onChangeState: setPassword
                    }
                ]}/>
            </SigninContent>
        </View>
    )
}