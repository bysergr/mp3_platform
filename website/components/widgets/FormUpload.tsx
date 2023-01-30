import axios from "axios"
import { useState } from "react"

const FormUpload = () => {
  const [content, setContent] = useState({
    name: '',
    file: ''
  })
  const [name, setName] = useState("Choose File")

  const handleChange = (e: any) => {
    if (e.target.name === 'file') {
      const show: string = e.target.value.split('\\').pop()
      if (!show) { 
        setName("Choose File")
      } else {
        setName(show)
      }


      setContent({
        ...content,
        file: e.target.files[0]
      })

      return
    }
  
    setContent({
      ...content,
      [e.target.name]: e.target.value
    })
  }


  const handleSubmit = async (e: any) => {
    e.preventDefault()
   
    if (!content.file) {
      alert("Don't File")
      return
    }

    const api = process.env.API_ADDRESS
    
    const data = new FormData()
    data.append("name", content.name)
    data.append("file", content.file)

    const res = axios({
      method: 'post',
      url: `${api}/upload`,
      headers: {'Content-Type': 'multipart/form-data'},
      data: data
    }) 
  }
  return (
    <form  onSubmit={handleSubmit} >

      <label htmlFor="name" className="my-2 mx-1 font-semibold block text-lg">Name</label>
      <input onChange={handleChange} required id="name" name="name" type="text" placeholder="name" className="border mb-6 bg-white border-gray-300 font-semibold block px-4 py-2 rounded w-full my-3" />

      <label htmlFor="file" className="my-6 w-full mx-auto border border-gray-300 rounded font-semibold block py-2 text-lg text-center">{name}</label>
      <input onChange={handleChange} required id="file" name="file" type="file" accept="audio/mp3" className="mb-6 w-full hidden" />

      <button className="mx-auto block font-semibold border border-zinc-600 w-40 py-2 mt-12 rounded">Send</button>
    </form>
  )
}

export default FormUpload
