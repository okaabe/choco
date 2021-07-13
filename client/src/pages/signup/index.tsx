import React from 'react';

import { Header } from '../../components/header';
import {
    SignUpContainer,
    SignUpContent,
    SignUpForm,
    SignUpFormInputs,
    SignUpThumbnail,
    SignUpButton,
    SignFormMessage
} from './styles';

import Thumbnail from '../../assets/imgs/thumb.svg'

export const SignUp: React.FC = () => {
    return (
        <SignUpContainer>
            <Header />
            <SignUpContent>
                <SignUpForm>
                    <SignFormMessage>
                        <h1>Sign Up</h1>
                        <p>By continuing, you agree to our User Agreement and Privacy Policy.</p>
                    </SignFormMessage>
                    <SignUpFormInputs>
                        <input type="text" name="username" id="username" placeholder="Username"/>
                        <input type="text" name="email" id="email" placeholder="Email"/>
                        <input type="password" name="password" id="password" placeholder="Password"/>
                        <input type="password" name="password" id="password2" placeholder="Repeat the password"/>
                    </SignUpFormInputs>
                    <SignUpButton>
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