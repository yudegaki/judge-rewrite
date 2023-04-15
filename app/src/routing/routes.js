import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Main from './containers/Main';
import Login from './containers/Login';
import Rankings from './containers/Rankings';
import PublicRoute from './components/PublicRoute';
import PrivateRoute from './components/PrivateRoute';

const Routes = ({ isAuthenticated }) => {
    return (
        <Switch>
            <Route exact path="/" component={Main} />
            <PublicRoute exact path="/login" component={Login} />
            <PrivateRoute exact path="/rankings" component={Rankings} isAuthenticated={isAuthenticated} />
            <Route render={() => <h1>Page not found</h1>} />
        </Switch>
    );
};

export default Routes;
