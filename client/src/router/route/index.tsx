import React from 'react';

import {
    Route as DefaultRoute,
    Redirect,
} from 'react-router-dom';

import { RouteProperties } from './route.props';
import { useSession } from '../../context/session';

export const Route: React.FC<RouteProperties> = ({
    children,
    isPrivate,
    ...rest
}) => {
    const session = useSession();

    session.setToken("kekekeke")

    if (isPrivate && !session.data) {
        return (
            <Redirect to="/signin"/>
        );
    }

    return (
        <DefaultRoute {...rest}>
            { children }
        </DefaultRoute>
    );
}
