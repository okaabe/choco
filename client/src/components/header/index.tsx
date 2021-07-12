import React from 'react';

import {
    HeaderButtons,
    HeaderContainer,
    HeaderLogo
} from './styles';

import ChocoLogo from '../../assets/imgs/chocolate.png'

export const Header: React.FC = () => {
    return (
        <HeaderContainer>
            <HeaderLogo>
                <img src={ ChocoLogo } alt="choco"/>
            </HeaderLogo>
            <HeaderButtons>

            </HeaderButtons>
        </HeaderContainer>
    )
}