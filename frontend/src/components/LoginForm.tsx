// components/LoginForm.tsx

import React from 'react';

const LoginForm: React.FC = () => {
  return (
    <form className="space-y-4">
      <label htmlFor="loginEmail" className="block text-black mb-1">Email</label>
      <input type="email" id="loginEmail" className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" placeholder="Enter your email" />
      
      <label htmlFor="loginPassword" className="block text-black mb-1">Password</label>
      <input type="password" id="loginPassword" className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:border-blue-500" placeholder="Enter your password" />
      
      <button type="submit" className="w-full bg-blue-500 text-black py-2 rounded-lg hover:bg-blue-600">Login</button>
    </form>
  );
};

export default LoginForm;
