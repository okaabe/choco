import React from 'react';

import { BrowserRouter } from 'react-router-dom';
import { Route } from './route';

import { SessionProvider } from '../context/session';

import { Home } from '../pages/home';

export const ApplicationRoutes: React.FC = () => {
    return (
        <SessionProvider>
            <BrowserRouter>
                <Route path="/" exact component={ Home }/>
            </BrowserRouter>
        </SessionProvider>
    )
}