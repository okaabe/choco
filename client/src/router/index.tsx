import React from 'react';

import { BrowserRouter } from 'react-router-dom';
import { Route } from './route';

import { toast } from 'react-toastify';

export const ApplicationRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Route path="/" exact component={() => {
                return (
                    <button onClick={() => toast("hello world")}>Hello World</button>
                )
            }}/>
        </BrowserRouter>
    )
}