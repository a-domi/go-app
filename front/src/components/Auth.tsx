import { useState, FormEvent } from "react"
import { useMutateAuth } from "../hooks/useMutateAuth"
import { CheckBadgeIcon, ArrowPathIcon } from '@heroicons/react/24/solid'

export const Auth = () => {
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [isLogin, setIsLogin] = useState(true)
	const { loginMutation, registerMutation } = useMutateAuth()

	const submitAuthHandler = async(e: FormEvent<HTMLFormElement>) => {
		e.preventDefault()
		if(isLogin){
			loginMutation.mutate({email: email, password: password})
		} else {
			await registerMutation.mutateAsync({
				email: email,
				password: password
			})
			.then(() =>
				loginMutation.mutate({email: email, password: password})
			) 
		}
	}
	return (
		<div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
			<div className="flex items-center">
				<CheckBadgeIcon className="h-8 w-8 mr-2 text-blue-500" />
				<span className="text-center text-3xl font-extrabold">
				Todo app
				</span>
			</div>
			<h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
			<form onSubmit={submitAuthHandler}>
				<div>
				<input
					className="mb-3 px-3 text-sm py-2 border border-gray-300"
					name="email"
					type="email"
					autoFocus
					placeholder="Email"
					onChange={(e) => setEmail(e.target.value)}
					value={email}
				/>
				</div>
				<div>
				<input
					className="mb-3 px-3 text-sm py-2 border border-gray-300"
					name="password"
					type="password"
					placeholder="Password"
					onChange={(e) => setPassword(e.target.value)}
					value={password}
				/>
				</div>
				<div className="flex justify-center my-2">
				<button
					className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
					disabled={!email || !password}
					type="submit"
				>
					{isLogin ? 'Login' : 'Sign Up'}
				</button>
				</div>
			</form>
			<ArrowPathIcon
				onClick={() => setIsLogin(!isLogin)}
				className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
			/>
			</div>
	)
}
