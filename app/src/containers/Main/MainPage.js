import React from 'react';
import { Switch, Route } from 'react-router-dom';
import HomePage from '../../components/HomePage/HomePage';
import AboutPage from '../../components/AboutPage/AboutPage';
import ContactPage from '../../components/ContactPage/ContactPage';

const MainPage = () => {
    return (
        <Switch>
            <Route exact path="/" component={HomePage} />
            <Route exact path="/about" component={AboutPage} />
            <Route exact path="/contact" component={ContactPage} />
        </Switch>
    );
};

export default MainPage;
