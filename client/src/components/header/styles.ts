import styled from 'styled-components';

export const HeaderContainer = styled.div`
    width: 100vw;
    height: 80px;

    background-color: #ffffff;

    display: flex;
    align-items: center;
    justify-content: space-between;
`;

export const HeaderLogo = styled.div`
    width: 50%;
    height: 100%;

    display: flex;
    align-items: center;

    padding: 0 15px;

    img {
        width: 60px;
        height: 60px;
    }
`;

export const HeaderButtons = styled.div`
    width: 50%;
    height: 100%;

    .signup, .signin, .profile {
       width: 300px;
       height: 60px;

       display: flex;
       align-items: center;
       justify-content: center;

       text-decoration: none;
       color: #fff; 
    }

    .signup {
        
    }

    .signin {

    }

    .profile {

    }
`;