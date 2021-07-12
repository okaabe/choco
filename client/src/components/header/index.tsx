import React from 'react';

import {
    HeaderButtons,
    HeaderContainer,
    HeaderLogo
} from './styles';

import {
    Link
} from 'react-router-dom';

import ChocoLogo from '../../assets/imgs/logo.svg'

export const Header: React.FC = () => {
    return (
        <HeaderContainer>
            <HeaderLogo>
                <Link to="/">
                    <img src={ ChocoLogo } alt="choco"/>
                </Link>
            </HeaderLogo>
            <HeaderButtons>

            </HeaderButtons>
        </HeaderContainer>
    )
}