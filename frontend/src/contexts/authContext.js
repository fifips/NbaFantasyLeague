import {createContext, useEffect, useState} from "react";


// AuthContext provides context for maintaining user login status
export const AuthContext = createContext({
    isLoggedIn: false,
    login: () => {},
    logout: () => {}
});

export const AuthProvider = ({ children }) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const login = (token) => {
        // Save the token to local storage
        sessionStorage.setItem('access_token', token);
        checkAuth();
    };

    const logout = () => {
        // Remove the token from local storage
        sessionStorage.removeItem('access_token');
        checkAuth();
    };

    // Check if the user is logged in on page load
    const checkAuth = async () => {
        // Get the token from local storage
        const token = sessionStorage.getItem('access_token');
        if (!token) {
            return setIsLoggedIn(false);
        }

        try {
            // Send the token to the backend to check if it's valid
            const res = await fetch('http://localhost:8080/user/is_logged_in', {
                credentials: "include",
            });

            // If the token is valid, set isLoggedIn to true
            if (res.statusCode === 200) {
                setIsLoggedIn(true);
            }
        } catch (err) {
            console.log('checkAuth request failed: ' + err);
        }
    };

    useEffect(() => {
        checkAuth();
    }, []);

    return (
        <AuthContext.Provider value={{isLoggedIn, login, logout}}>
            {children}
        </AuthContext.Provider>
    );
}