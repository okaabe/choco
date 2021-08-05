import React from 'react';

import { BrowserRouter } from 'react-router-dom';
import { Route } from './route';

import { SessionProvider } from '../context/session';

import { Home } from '../pages/home';
import { SignIn } from '../pages/session/signin';

export const ApplicationRoutes: React.FC = () => {
    return (
        <SessionProvider>
            <BrowserRouter>
                <Route path="/" exact component={ Home }/>
                <Route path="/signin" exact component={ SignIn } />
            </BrowserRouter>
        </SessionProvider>
    )
}