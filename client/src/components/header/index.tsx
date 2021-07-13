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
        <Link to="/profile" className="profile">Profile</Link>
    )
}

const SignInAndSignUpButtons: React.FC = () => {
    return (
        <>
            <Link to="/signup" className="signup">Sign Up</Link>
            <Link to="/signin" className="signin">Sign In</Link>
        </>
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