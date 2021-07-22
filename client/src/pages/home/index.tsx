import React from 'react';

import {
    HomeContainer,
    HomeContent
} from './styles';

import { Header } from '../../components/header';

import { useAuth } from '../../hooks/auth';

export const Home: React.FC = () => {
    const [ user ] = useAuth()

    return (
        <HomeContainer>
            <Header />
            <HomeContent>
                <h1>Hello { user.username }</h1>
            </HomeContent>
        </HomeContainer>
    );
}
