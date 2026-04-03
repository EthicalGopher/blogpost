import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../utils/api";
import { useAuth } from "../context/AuthContext";
import { Save, ArrowLeft, Eye, Sparkles } from "lucide-react";

const CreatePost: React.FC = () => {
  const { postId } = useParams();
  const { userId } = useAuth();
  const navigate = useNavigate();
  const isEditing = !!postId;

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(isEditing);

  useEffect(() => {
    if (isEditing) {
      api.get(`post/${postId}`).then((res) => {
        setTitle(res.data.title);
        setContent(res.data.content);
        setLoading(false);
      });
    }
  }, [postId]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const data = { title, content };
    try {
      if (isEditing) { await api.put(`user/${userId}/post/${postId}`, data); }
      else { await api.post(`user/${userId}/post`, data); }
      navigate(`/profile/${userId}`);
    } catch (err) { console.error(err); }
  };

  if (loading) return <div className="min-h-screen flex items-center justify-center font-black uppercase tracking-widest">Loading Editor...</div>;

  return (
    <div className="min-h-screen bg-white">
      <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-100 h-20 flex items-center px-4">
        <div className="container mx-auto flex justify-between items-center">
          <button onClick={() => navigate(-1)} className="p-2 hover:bg-slate-100 rounded-xl transition-colors text-slate-500"><ArrowLeft size={24} /></button>
          <button onClick={handleSubmit} className="btn-vibrant flex items-center gap-2 py-2.5"><Sparkles size={18} /><span>{isEditing ? "Update" : "Publish"} Story</span></button>
        </div>
      </nav>
      <main className="container mx-auto px-4 py-12 max-w-3xl">
        <input className="w-full text-5xl font-black placeholder:text-slate-200 focus:outline-none bg-transparent mb-8" placeholder="Title..." value={title} onChange={(e) => setTitle(e.target.value)} autoFocus />
        <textarea className="w-full text-xl placeholder:text-slate-300 focus:outline-none bg-transparent min-h-[500px] resize-none leading-relaxed text-slate-700" placeholder="Tell your story..." value={content} onChange={(e) => setContent(e.target.value)} />
      </main>
    </div>
  );
};

export default CreatePost;
