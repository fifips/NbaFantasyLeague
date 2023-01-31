import '../../styles/header/playerInformation.css';
import {useContext} from "react";
import {AuthContext} from "../../contexts/authContext";
import {Link} from "react-router-dom";
import {BouncingButton} from "../common/bouncingButton";


const LoggedInProfile = () => {
    const {logout} = useContext(AuthContext);

    return (
        <div className="logged-in-profile">
            <BouncingButton onClick={logout}>
                Logout
            </BouncingButton>
        </div>
    );
}

const LoggedOutProfile = () => {
    return (
        <div className="logged-out-profile">
            <Link to="/login">
                <BouncingButton>
                    Login
                </BouncingButton>
            </Link>
            <Link to="/register">
                <BouncingButton>
                    Register
                </BouncingButton>
            </Link>
        </div>
    );
}

export const UserInformation = () => {
    const {isLoggedIn} = useContext(AuthContext);

    return (
        <div className="user-profile">
            { isLoggedIn ? <LoggedInProfile /> : <LoggedOutProfile /> }
        </div>
    );
}

