import './styles/App.css';
import {Header} from './components/header/header'
import {BrowserRouter, Outlet, Route, Routes} from "react-router-dom";
import {Register} from "./components/register";
import {Activate} from "./components/activate";
import {Notifications} from "./components/common/notification";
import {useContext} from "react";
import {Login} from "./components/login";
import {NavigationMenu} from "./components/nagivationMenu/navigationMenu";
import {Leagues} from "./components/leagues";
import {AuthContext} from "./contexts/authContext";
import {Home} from "./components/home";

const App = () => {
    const {isLoggedIn} = useContext(AuthContext);

    return (
      <div className="App">
        <Notifications/>
        <BrowserRouter>
            <Header />
            <NavigationMenu />
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route element={ isLoggedIn ? <Outlet/> : <Login/>}>
                    <Route path="/leagues" element={<Leagues/>}/>
                </Route>
                <Route element={ !isLoggedIn ? <Outlet/> : <Home/>}>
                    <Route path="/login" element={<Login/>}/>
                    <Route path="/register" element={<Register/>}/>
                    <Route path="/user/activate" element={<Activate/>}/>
                </Route>
            </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;