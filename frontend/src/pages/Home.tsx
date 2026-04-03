import React, { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import api from "../utils/api";
import {
  LogOut,
  PlusCircle,
  BookOpen,
  User as UserIcon,
  TrendingUp,
  Sparkles,
  Search,
} from "lucide-react";
import { useNavigate } from "react-router-dom";

interface Post {
  ID: number;
  title: string;
  content: string;
  CreatedAt: string;
  author: {
    name: string;
  };
}

const Home: React.FC = () => {
  const { logout, userId } = useAuth();
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await api.get("post");
        setPosts(response.data || []);
        console.log(response.data);
      } catch (err) {
        console.error("Failed to fetch posts", err);
      } finally {
        setLoading(false);
      }
    };
    fetchPosts();
  }, []);

  return (
    <div className="min-h-screen bg-slate-50/50">
      <nav className="sticky top-0 z-50 bg-white/70 backdrop-blur-xl border-b border-slate-200">
        <div className="container mx-auto px-4 h-20 flex items-center justify-between">
          <div
            className="flex items-center gap-3 group cursor-pointer"
            onClick={() => navigate("/")}
          >
            <div className="w-10 h-10 bg-gradient-to-tr from-primary-600 to-accent-violet rounded-xl flex items-center justify-center text-white shadow-lg shadow-primary-500/20 group-hover:scale-110 transition-transform">
              <BookOpen size={24} />
            </div>
            <span className="text-2xl font-black tracking-tighter gradient-text">
              BlogSphere
            </span>
          </div>

          <div className="flex items-center gap-4">
            <button
              onClick={() => navigate("/create-post")}
              className="btn-vibrant flex items-center gap-2"
            >
              <PlusCircle size={20} />
              <span className="hidden sm:inline">New Story</span>
            </button>
            <div className="h-10 w-px bg-slate-200 mx-2 hidden sm:block"></div>
            <button
              className="p-2.5 rounded-xl border border-slate-200 hover:bg-slate-50 transition-colors text-slate-600"
              onClick={() => navigate(`/profile/${userId}`)}
            >
              <UserIcon size={22} />
            </button>
            <button
              className="p-2.5 rounded-xl text-accent-pink hover:bg-accent-pink/5 transition-colors"
              onClick={logout}
            >
              <LogOut size={22} />
            </button>
          </div>
        </div>
      </nav>

      <main className="container mx-auto px-4 pb-20">
        <header className="relative py-20 overflow-hidden text-center">
          <h1 className="text-5xl md:text-7xl font-black mb-6 tracking-tight">
            Write your <span className="gradient-text">masterpiece.</span>
          </h1>
          <p className="text-xl text-slate-600 max-w-2xl mx-auto mb-10">
            Join a community of forward-thinking writers and share your unique
            perspective.
          </p>
        </header>

        {loading ? (
          <div className="text-center py-20 text-slate-400 font-bold uppercase tracking-widest">
            Loading latest stories...
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {posts.map((post, idx) => (
              <article
                key={post.ID}
                className="glass-card group hover:-translate-y-2 transition-all duration-500 flex flex-col h-full"
              >
                <div className="p-8 flex flex-col flex-grow">
                  <div className="flex items-center gap-2 mb-4 text-slate-400 text-xs font-bold uppercase tracking-wider">
                    <span>{new Date(post.CreatedAt).toLocaleDateString()}</span>
                  </div>
                  <h3 className="text-2xl font-bold mb-4 line-clamp-2 group-hover:text-primary-600 transition-colors">
                    {post.title}
                  </h3>
                  <p className="text-slate-600 text-sm mb-8 line-clamp-3 leading-relaxed">
                    {post.content}
                  </p>
                  <div className="mt-auto pt-6 border-t border-slate-100 flex items-center justify-between">
                    <div className="flex items-center gap-3">
                      <div className="w-10 h-10 rounded-xl bg-slate-100 flex items-center justify-center text-slate-500 ring-2 ring-white">
                        <UserIcon size={18} />
                      </div>
                      <span className="text-sm font-bold text-slate-700">
                        {post.author?.name || "Anonymous"}
                      </span>
                    </div>
                    <button
                      onClick={() => navigate(`/post/${post.ID}`)}
                      className="text-primary-600 font-black text-sm hover:translate-x-1 transition-transform"
                    >
                      READ →
                    </button>
                  </div>
                </div>
              </article>
            ))}
          </div>
        )}
      </main>
    </div>
  );
};

export default Home;
