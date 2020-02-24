import React from 'react';
import { BrowserRouter } from "react-router-dom";
import Theme from './containers/Theme';
import Router from "./containers/Router";

function App() {
    return (
        <Theme>
            <BrowserRouter>
                <Router />
            </BrowserRouter>
        </Theme>
	);
}

export default App;
