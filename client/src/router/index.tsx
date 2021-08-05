import React from 'react';

import { BrowserRouter } from 'react-router-dom';
import { Route } from './route';

export const ApplicationRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Route path="/"/>
        </BrowserRouter>
    )
}