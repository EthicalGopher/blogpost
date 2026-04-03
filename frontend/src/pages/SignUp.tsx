import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import api from "../utils/api";
import { Mail, Lock, User, AlertCircle, BookOpen, CheckCircle, ArrowRight } from "lucide-react";

const SignUp: React.FC = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      await api.post("signup", { name, email, password });
      setSuccess(true);
      setTimeout(() => navigate("/signin"), 2000);
    } catch (err: any) {
      setError(err.response?.data?.error || "Registration failed.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row-reverse">
      <div className="hidden md:flex md:w-1/2 bg-gradient-to-bl from-accent-violet via-primary-600 to-accent-emerald items-center justify-center p-12 text-white">
        <div className="max-w-md relative z-10 animate-fade-in-up">
          <h1 className="text-6xl font-black mb-6 leading-tight tracking-tighter text-right">Start your legacy.</h1>
          <p className="text-xl text-white/80 font-medium text-right">Join a platform built for creators who demand the best.</p>
        </div>
      </div>
      <div className="flex-1 flex items-center justify-center p-8 bg-slate-50/50">
        <div className="w-full max-w-md">
          <h2 className="text-4xl font-black mb-2 text-slate-900">Get Started.</h2>
          <p className="text-slate-500 font-bold mb-10">Join the elite community of creators</p>
          {error && <div className="bg-accent-pink/10 text-accent-pink p-4 rounded-2xl mb-8 flex items-center gap-3 font-bold text-sm"><AlertCircle size={20} />{error}</div>}
          {success && <div className="bg-emerald-500/10 text-emerald-600 p-4 rounded-2xl mb-8 flex items-center gap-3 font-bold text-sm"><CheckCircle size={20} />Registration successful!</div>}
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="relative"><User className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} /><input type="text" className="input-vibrant pl-12" placeholder="Full Name" value={name} onChange={(e) => setName(e.target.value)} required /></div>
            <div className="relative"><Mail className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} /><input type="email" className="input-vibrant pl-12" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required /></div>
            <div className="relative"><Lock className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400" size={20} /><input type="password" className="input-vibrant pl-12" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} required /></div>
            <button type="submit" className="btn-vibrant w-full py-4 flex items-center justify-center gap-2" disabled={loading || success}><span className="text-lg uppercase tracking-widest">{loading ? "Processing..." : "Create Account"}</span><ArrowRight size={20} /></button>
          </form>
          <div className="mt-12 text-center border-t border-slate-200 pt-8"><p className="text-slate-500 font-bold">Already a member? <Link to="/signin" className="text-primary-600 underline underline-offset-4 decoration-2">Sign In</Link></p></div>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
