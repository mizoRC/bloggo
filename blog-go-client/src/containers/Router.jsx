import React from "react";
import { Switch, Route, withRouter, Redirect } from "react-router-dom";
import Landing from '../components/Landing';
import Test from '../components/Test';

const Router = () => {
    const redirectToRoot = () => {
		return <Redirect to="/" />
	};

	return (
		<Switch>
            <Route exact path="/" component={Landing} />
            <Route exact path="/test" component={Test} />

            {/* Mantener como ultima para que sepa cuales son las rutas que si son validas */}
            <Route path="*" component={redirectToRoot} />
		</Switch>
	);
};

export default withRouter(Router);