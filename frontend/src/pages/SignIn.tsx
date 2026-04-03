import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import api from "../utils/api";
import { LogIn, Mail, Lock, AlertCircle, BookOpen, ArrowRight, Sparkles } from "lucide-react";

const SignIn: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      const response = await api.post("signin", { email, password });
      const { user_id, access_token, refresh_token } = response.data;
      login(user_id.toString(), access_token, refresh_token);
      navigate("/");
    } catch (err: any) {
      setError(err.response?.status === 401 ? "Invalid credentials" : "An error occurred");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row">
      <div className="hidden md:flex md:w-1/2 bg-gradient-to-br from-primary-600 via-accent-violet to-accent-pink items-center justify-center p-12 text-white relative">
        <div className="max-w-md relative z-10 animate-fade-in-up">
          <BookOpen size={40} className="mb-8" />
          <h1 className="text-6xl font-black mb-6 leading-tight tracking-tighter">Your words matter.</h1>
          <p className="text-xl text-white/80 font-medium">Join thousands of creators sharing their insights with a global audience.</p>
        </div>
      </div>
      <div className="flex-1 flex items-center justify-center p-8 bg-slate-50/50">
        <div className="w-full max-w-md">
          <h2 className="text-4xl font-black mb-2 text-slate-900">Welcome Back.</h2>
          <p className="text-slate-500 font-bold mb-10">Continue your creative journey</p>
          {error && <div className="bg-accent-pink/10 border border-accent-pink/20 text-accent-pink p-4 rounded-2xl mb-8 flex items-center gap-3"><AlertCircle size={20} /><span className="text-sm font-bold">{error}</span></div>}
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="relative"><Mail className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} /><input type="email" className="input-vibrant pl-12" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required /></div>
            <div className="relative"><Lock className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} /><input type="password" className="input-vibrant pl-12" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} required /></div>
            <button type="submit" className="btn-vibrant w-full py-4 flex items-center justify-center gap-2 group" disabled={loading}><span className="text-lg uppercase tracking-widest">{loading ? "Signing in..." : "Sign In"}</span><ArrowRight size={20} /></button>
          </form>
          <div className="mt-12 text-center border-t border-slate-200 pt-8"><p className="text-slate-500 font-bold">New to BlogSphere? <Link to="/signup" className="text-primary-600 underline underline-offset-4 decoration-2">Create an account</Link></p></div>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
