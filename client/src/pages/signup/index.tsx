import React from 'react';

import { Header } from '../../components/header';
import {
    SignUpContainer,
    SignUpContent,
    SignUpForm,
    SignUpFormInputs,
    SignUpThumbnail,
    SignUpButton,
    SignFormMessage,
    SignUpErrorMessage
} from './styles';

import { toast } from 'react-toastify';

import Thumbnail from '../../assets/imgs/thumb.svg'

export const SignUp: React.FC = () => {
    const [username, setUsername] = React.useState<string>();
    const [email, setEmail] = React.useState<string>();
    const [password, setPassword] = React.useState<string>();
    const [confirmPassword, setConfirmPassword] = React.useState<string>();
    const [err, setErr] = React.useState<string>();

    const onSubmit = React.useCallback((e) => {
        e.preventDefault();

        if (
            !username ||
            !email ||
            !password ||
            !confirmPassword
        ) {
            setErr("Fill all the fields.")
        }
    }, [username, email, password, confirmPassword])

    return (
        <SignUpContainer>
            <Header />
            <SignUpContent>
                <SignUpForm onSubmit={ onSubmit }>
                    <SignFormMessage>
                        <h1>Sign Up</h1>
                        <p>By continuing, you agree to our User Agreement and Privacy Policy.</p>
                    </SignFormMessage>
                    <SignUpFormInputs>
                        <input type="text" name="username" id="username" placeholder="Username" onChange={(e) => setUsername(e.target.value)}/>
                        <input type="text" name="email" id="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)}/>
                        <input type="password" name="password" id="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
                        <input type="password" name="password" id="confirmPassword" placeholder="Repeat the password" onChange={(e) => setConfirmPassword(e.target.value)}/>
                    </SignUpFormInputs>
                    <SignUpErrorMessage>{ err }</SignUpErrorMessage>
                    <SignUpButton type="submit">
                        Sign Up
                    </SignUpButton>
                </SignUpForm>
                <SignUpThumbnail>
                    <img src={ Thumbnail } alt=""/>
                </SignUpThumbnail>
            </SignUpContent>
        </SignUpContainer>
    )
}