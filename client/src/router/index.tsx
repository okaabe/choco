import React from 'react';

import { BrowserRouter } from 'react-router-dom';
import { Route } from './route';

import { Home } from '../pages/home';

export const ApplicationRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Route path="/" exact component={ Home }/>
        </BrowserRouter>
    )
}