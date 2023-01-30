import './styles/App.css';
import {Header} from './components/header/header'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Register} from "./components/register";
import {Activate} from "./components/activate";
import {Notification, Notifications} from "./components/common/notification";
import {useContext} from "react";
import {NotificationContext} from "./contexts/notifContext";

const App = () => {
    return (
      <div className="App">
        <Notifications/>
        <BrowserRouter>
            <Header />
            <Routes>
                <Route path="/register" element={<Register/>}/>
                <Route path="/user/activate" element={<Activate/>}/>
            </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;