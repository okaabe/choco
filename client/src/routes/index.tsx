import React from 'react';

import { BrowserRouter, Switch } from 'react-router-dom';
import { Route } from './route';

import { SignUp } from '../pages/signup';
import { SignIn } from '../pages/signin';
import { Home } from '../pages/home';

export const AppRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact isPrivate component={ Home }/>
                <Route path="/signin" component={ SignIn }/>
                <Route path="/signup" component={ SignUp }/>
            </Switch>
        </BrowserRouter>
    )
}