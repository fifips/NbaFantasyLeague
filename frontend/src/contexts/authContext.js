import {createContext, useEffect, useState} from "react";
import {isUserLoggedIn, loginUser, logoutUser} from "../api/user_api";

// AuthContext provides context for maintaining user login status
export const AuthContext = createContext({
    isLoggedIn: false,
    login: () => {},
    logout: () => {},
});

export const AuthProvider = ({ children }) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const login = async ({email, password, rememberMe}) => {
        const r = await loginUser({email, password, rememberMe})
        setIsLoggedIn(r.ok)
        return r.ok
    };

    const logout = () => {
        logoutUser().then(r => {
            if (r.ok) setIsLoggedIn(false)
            return r.ok
        })
    };

    useEffect(() => {
        isUserLoggedIn().then((r) => {
            setIsLoggedIn(r.ok)
            });
    },[])

    useEffect(() => {
        sessionStorage.setItem("isLoggedIn", String(isLoggedIn))
    },[isLoggedIn])

    return (
        <AuthContext.Provider value={{isLoggedIn, login, logout}}>
            {children}
        </AuthContext.Provider>
    );
}