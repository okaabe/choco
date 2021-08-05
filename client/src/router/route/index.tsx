import React from 'react';

import {
    Route as DefaultRoute,
    Redirect,
} from 'react-router-dom';

import { RouteProperties } from './route.props';


export const Route: React.FC<RouteProperties> = ({
    children,
    isPrivate,
    ...rest
}) => {
    if (isPrivate) {
        return (
            <Redirect to="/signin"/>
        );
    }

    return (
        <DefaultRoute>
            { children }
        </DefaultRoute>
    );
}
