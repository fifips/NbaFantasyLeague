import '../../styles/header/playerInformation.css';
import {useContext} from "react";
import {AuthContext} from "../../contexts/authContext";
import {Link} from "react-router-dom";
import {registerUrl} from "../../const";
import {BouncingButton} from "../common/bouncingButton";


const LoggedInProfile = ({logout}) => {
    return (
        <div className="logged-in-profile">
            <BouncingButton onClick={logout}>
                Logout
            </BouncingButton>
        </div>
    );
}

const LoggedOutProfile = ({login}) => {
    return (
        <div className="logged-out-profile">
            <BouncingButton onClick={login}>
                Login
            </BouncingButton>
            <Link to="/register">
                <BouncingButton>
                    Register
                </BouncingButton>
            </Link>
        </div>
    );
}

export const UserInformation = () => {
    const {isLoggedIn, login, logout} = useContext(AuthContext);

    return (
        <div className="user-profile">
            {
                isLoggedIn
                    ? <LoggedInProfile logout={logout}/>
                    : <LoggedOutProfile login={login}/>
            }
        </div>
    );
}

