import axios from "axios"
import { useState } from "react"

const FormLogin = () => {
  const [credentials, setCredentials] = useState({
    email: '',
    password: ''
  })

  const handleChange = (e: any) => {
    setCredentials({
      ...credentials,
      [e.target.name]: e.target.value
    })
    
  }
  
  const handleSubmit = async (e: any) => {
    const api = process.env.API_ADDRESS

    e.preventDefault()
    const res = await axios({ 
      method: 'post',
      url: `${api}/login`,
      auth: {
        username: credentials['email'],
        password: credentials['password'],
      }
    })
    
    console.log(res.data)
  }

  return (
    <form onSubmit={handleSubmit}>
      <label htmlFor="email" className="my-2 mx-1 font-semibold block text-lg">Email</label>
      <input onChange={handleChange} required id="email"  autoComplete="username" name="email" type="email" placeholder="Email" className="border mb-6 bg-white font-semibold block px-4 py-2 rounded w-full my-3" />

      <label htmlFor="password" className="my-2 mx-1 font-semibold block text-lg">Password</label>
      <input onChange={handleChange} required id="password" autoComplete="current-password" name="password" type="password" placeholder="Password" className="border mb-6  bg-white lock px-4 py-2 rounded w-full my-3" />

      <button className="mx-auto block font-semibold border border-zinc-600 w-40 py-2 rounded">Login</button>
    </form>
  )
}

export default FormLogin
