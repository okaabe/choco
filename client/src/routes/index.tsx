import React from 'react';

import { BrowserRouter, Switch } from 'react-router-dom';
import { Route } from './route';

import { SignUp } from '../pages/signup';

export const AppRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact isPrivate component={() => {
                    return (
                        <h1>Pagina privada onde apenas pessoas autenticadas podem acessar.</h1>
                    )
                }}/>
                <Route path="/signup" component={ SignUp }/>
            </Switch>
        </BrowserRouter>
    )
}