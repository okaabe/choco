import React from 'react';

import {
    HeaderButtons,
    HeaderContainer,
    HeaderLogo
} from './styles';

import {
    Link
} from 'react-router-dom';

import { useAuth } from '../../hooks/auth';

import ChocoLogo from '../../assets/imgs/logo.svg'

const ProfileButton: React.FC<{ username: string; }> = ({
    username
}) => {
    return (
        <h1>{username}</h1>
    )
}

const SignInAndSignUpButtons: React.FC = () => {
    return (
        <h1>Hello World</h1>
    )
}

export const Header: React.FC = () => {
    const [user] = useAuth()

    return (
        <HeaderContainer>
            <HeaderLogo>
                <Link to="/">
                    <img src={ ChocoLogo } alt="choco"/>
                </Link>
            </HeaderLogo>
            <HeaderButtons>
                {
                    user.username && user.jwt ?
                    <ProfileButton username={ user.username }/> :
                    <SignInAndSignUpButtons />
                }
            </HeaderButtons>
        </HeaderContainer>
    )
}