import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../utils/api";
import { useAuth } from "../context/AuthContext";
import {
  User as UserIcon,
  Mail,
  BookOpen,
  Trash2,
  Edit,
  ArrowLeft,
} from "lucide-react";

interface User {
  ID: number;
  name: string;
  email: string;
}

interface Post {
  ID: number;
  title: string;
  content: string;
  CreatedAt: string;
}

const Profile: React.FC = () => {
  const { id } = useParams();
  const { userId, logout } = useAuth();
  const [user, setUser] = useState<User | null>(null);
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const isOwnProfile = Number(userId) === Number(id);

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const userRes = await api.get(`user/${id}/`);
        setUser(userRes.data);
        const postsRes = await api.get(`user/${id}/post`);
        setPosts(postsRes.data || []);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };
    fetchProfile();
  }, [id]);

  const handleDeletePost = async (postId: number) => {
    if (!window.confirm("Delete permanently?")) return;
    try {
      await api.delete(`user/${userId}/post/${postId}`);
      setPosts(posts.filter((p) => p.id !== postId));
    } catch (err) {
      console.error(err);
    }
  };

  if (loading)
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      </div>
    );
  if (!user)
    return (
      <div className="text-center py-20 text-2xl font-black">
        Creator not found.
      </div>
    );

  return (
    <div className="min-h-screen bg-slate-50/50">
      <nav className="bg-white border-b border-slate-200 py-4">
        <div className="container mx-auto px-4">
          <button
            onClick={() => navigate("/")}
            className="flex items-center gap-2 text-slate-500 hover:text-primary-600 font-bold transition-colors"
          >
            <ArrowLeft size={20} />
            Back
          </button>
        </div>
      </nav>
      <header className="bg-white border-b border-slate-200 py-12 mb-8">
        <div className="container mx-auto px-4 flex items-center gap-8">
          <div className="w-32 h-32 rounded-3xl bg-gradient-to-tr from-primary-600 to-accent-violet flex items-center justify-center text-white shadow-xl">
            <UserIcon size={64} />
          </div>
          <div>
            <h1 className="text-4xl font-black mb-2">{user.name}</h1>
            <div className="flex gap-6 text-slate-500 font-bold text-sm uppercase tracking-wider">
              <span>
                <Mail size={16} className="inline mr-2" />
                {user.email}
              </span>
              <span>
                <BookOpen size={16} className="inline mr-2" />
                {posts.length} Stories
              </span>
            </div>
            {isOwnProfile && (
              <button
                onClick={logout}
                className="mt-4 px-4 py-1.5 rounded-lg border border-slate-200 text-sm font-bold text-slate-500 hover:bg-slate-50 transition-all"
              >
                Sign Out
              </button>
            )}
          </div>
        </div>
      </header>
      <main className="container mx-auto px-4 pb-20">
        <h2 className="text-2xl font-black mb-8">
          {isOwnProfile ? "Your Work" : "Published work"}
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {posts.map((post) => (
            <div
              key={post.ID}
              className="glass-card p-8 group transition-all duration-300"
            >
              <div className="flex justify-between mb-4">
                <span className="text-[10px] font-black uppercase tracking-[0.2em] text-primary-500">
                  Story
                </span>
                {isOwnProfile && (
                  <div className="flex gap-2">
                    <button
                      onClick={() => navigate(`/edit-post/${post.ID}`)}
                      className="p-2 bg-slate-50 rounded-lg text-slate-400 hover:text-primary-600 transition-colors"
                    >
                      <Edit size={16} />
                    </button>
                    <button
                      onClick={() => handleDeletePost(post.ID)}
                      className="p-2 bg-slate-50 rounded-lg text-slate-400 hover:text-accent-pink transition-colors"
                    >
                      <Trash2 size={16} />
                    </button>
                  </div>
                )}
              </div>
              <h3 className="text-2xl font-bold mb-4">{post.title}</h3>
              <p className="text-slate-600 line-clamp-2 mb-6">{post.content}</p>
              <div className="flex items-center justify-between pt-6 border-t border-slate-100">
                <span className="text-xs font-bold text-slate-400">
                  {new Date(post.CreatedAt).toLocaleDateString()}
                </span>
                <button
                  onClick={() => navigate(`/post/${post.ID}`)}
                  className="text-primary-600 font-black text-xs uppercase tracking-widest"
                >
                  Read →
                </button>
              </div>
            </div>
          ))}
        </div>
      </main>
    </div>
  );
};

export default Profile;
