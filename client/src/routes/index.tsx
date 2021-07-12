import React from 'react';

import { BrowserRouter, Switch } from 'react-router-dom';
import { Route } from './route';

export const AppRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact isPrivate component={() => {
                    return (
                        <h1>Hello World</h1>
                    )
                }}/>
                <Route path="/signin" component={() => {
                    return (
                        <h1>Sign in mano</h1>
                    )
                }}/>
            </Switch>
        </BrowserRouter>
    )
}