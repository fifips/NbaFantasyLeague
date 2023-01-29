import '../../styles/header/header.css';
import {Schedule} from "./schedule";
import {UserInformation} from "./playerInformation";

export const Header = () => {
    return (
        <div className="header">
            <UserInformation />
            <Schedule />
        </div>
    );
}