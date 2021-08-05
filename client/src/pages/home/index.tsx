import React from 'react';

import { View } from '../../components/view';
import { toast } from 'react-toastify';

export const Home: React.FC = () => {
    React.useEffect(() => {
        toast("Private session")
    }, [])

    return (
        <View>
            Hello Okaabe, you are on the home page. be careful
        </View>
    );
}
