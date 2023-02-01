import {BouncingButton} from "./common/bouncingButton";
import {useContext, useState} from "react";
import "../styles/login.css"
import {NotificationContext} from "../contexts/notifContext";
import {Navigate} from "react-router-dom";
import {AuthContext} from "../contexts/authContext";

export const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loginCompleted, setLoginCompleted] = useState(false);
    const [rememberMe, setRememberMe] = useState(false);
    const { addNotification } = useContext(NotificationContext);
    const { isLoggedIn, login } = useContext(AuthContext);

    const handleSubmit = event => {
        event.preventDefault();

        if (!email || !password) {
            addNotification("Please fill in all the fields");
            return;
        }
        if (!/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(email)) {
            addNotification("Please enter a valid email address");
            return;
        }
        if (password.length < 8) {
            addNotification("Password should be at least 8 characters long");
            return;
        }

        login({email, password, rememberMe}) ? setLoginCompleted(true) : addNotification("Wrong e-mail or password")
    }

    if (loginCompleted || isLoggedIn) return <Navigate to={"/"}/>
    return (
        <div className="login-page">
            <form className="login" onSubmit={handleSubmit}>
                <input type="email" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} />
                <input type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} />
                <label className="remember-me">
                    <input className="remember-me" type="checkbox" onChange={() => setRememberMe(!rememberMe)}/>
                    Remember me
                </label>
                <BouncingButton className="login" type="submit">Login</BouncingButton>
            </form>
        </div>
    );
}