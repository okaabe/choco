import React from 'react';

import {
    Route as BaseRoute,
    RouteProps as BaseRouteProps,
    Redirect,
} from 'react-router-dom';

import { useAuth } from '../hooks/auth'

export interface RouteProps extends BaseRouteProps {
    isPrivate?: boolean;
}

export const Route: React.FC<RouteProps> = ({
    isPrivate,
    ...baseProps
}) => {
    const [ user ] = useAuth()

    if (isPrivate && !(user && user.jwt)) {
        return <Redirect to="/signup"/>
    }

    return (
        <BaseRoute {...baseProps} />        
    )
}
