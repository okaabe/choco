import React from 'react';
import { Redirect } from 'react-router-dom';

import { Header } from '../../components/header';
import {
    SignInContainer,
    SignInContent,
    SignInForm,
    SignInFormInputs,
    SignInThumbnail,
    SignInButton,
    SignFormMessage,
    SignInErrorMessage
} from './styles';

import Thumbnail from '../../assets/imgs/thumb.svg';

import { useAuth } from '../../hooks/auth';
import { signin } from '../../services/auth';

export const SignIn: React.FC = () => {
    const [email, setEmail] = React.useState<string>();
    const [password, setPassword] = React.useState<string>();
    const [err, setErr] = React.useState<string>();
    const [user, setters] = useAuth();
    const onSubmit = React.useCallback((e) => {
        e.preventDefault();

        if (
            !email ||
            !password
        ) {
            return setErr("Fill all the fields.")
        }

        signin(email, password)
            .then((user) => {
                if (!user) {
                    console.log(user);
                    return setErr("Something weird happened, check the console.")
                }
                
                setters.setUser!(user)
                setters.setToken!(user.jwt!)
            }).catch((signUpErr) => {
                console.log(signUpErr);
                setErr(String(signUpErr));
            })
    }, [email, password, setters])

    if (user && user.jwt) {
        return <Redirect to="/"/>
    }


    return (
        <SignInContainer>
            <Header />
            <SignInContent>
                <SignInForm onSubmit={ onSubmit }>
                    <SignFormMessage>
                        <h1>Sign In</h1>
                        <p>By continuing, you agree to our User Agreement and Privacy Policy.</p>
                    </SignFormMessage>
                    <SignInFormInputs>
                        <input type="text" name="email" id="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)}/>
                        <input type="password" name="password" id="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
                    </SignInFormInputs>
                    <SignInErrorMessage>{ err }</SignInErrorMessage>
                    <SignInButton type="submit">
                        Sign In
                    </SignInButton>
                </SignInForm>
                <SignInThumbnail/>
            </SignInContent>
        </SignInContainer>
    )
}