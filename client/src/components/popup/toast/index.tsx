import React from 'react';

import {
    ToastContainer,
    toast
} from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css'

export const Toast: React.FC = () => {
    toast.configure();

    return <ToastContainer />;
}