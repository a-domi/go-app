import React from 'react';
import axios from "axios";
import { useState,useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import { CsrfToken } from './types'

function App() {
	useEffect(() => {
		axios.defaults.withCredentials = true
		const getCsrfToken = async() => {
			const { data } = await axios.get<CsrfToken>(
				'http://localhost:8080/csrf'
			)
			axios.defaults.headers.common['X-CSRF-TOKEN'] = data.csrf_token
		}
		getCsrfToken()
	})
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Auth />} />
				<Route path="todo" element={<Todo />} />
			</Routes>
		</BrowserRouter>
	);
}

export default App;
