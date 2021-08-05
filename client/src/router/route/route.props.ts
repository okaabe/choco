import {
    RouteProps as DefaultRouteProps
} from 'react-router-dom';

export type RouteProperties = DefaultRouteProps & {
    isPrivate?: boolean;
}
