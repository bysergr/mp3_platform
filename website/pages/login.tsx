import FormLogin from "@/components/widgets/FormLogin";

const Login = () => {
  return (
    <div className="max-w-md w-full bg-white block rounded-lg my-10 p-6 px-10 border border-gray-300">
      <h2 className="font-semibold text-2xl text-center my-6">SIGN IN</h2>
      <FormLogin />
    </div>
  );
};

export default Login;
