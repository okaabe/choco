import React from 'react';

import { ViewContainer } from './view.styles';

import { ViewProperties } from './view.props';

export const View: React.FC<ViewProperties> = ({
    allowedContextMenu,
    children
}) => {
    return (
        <ViewContainer>
            { children }
        </ViewContainer>
    )
}
