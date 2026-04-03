import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../utils/api";
import { useAuth } from "../context/AuthContext";
import {
  MessageSquare,
  Send,
  Trash2,
  ArrowLeft,
  User as UserIcon,
  Calendar,
  Clock,
  Share2,
  Bookmark,
  Sparkles,
} from "lucide-react";

interface Comment {
  id: number;
  text: string;
  CreatedAt: string;
  user_id: number;
  user: { name: string };
}

interface Post {
  id: number;
  title: string;
  content: string;
  CreatedAt: string;
  author: { id: number; name: string };
}

const PostDetail: React.FC = () => {
  const { postId } = useParams();
  const { userId } = useAuth();
  const [post, setPost] = useState<Post | null>(null);
  const [comments, setComments] = useState<Comment[]>([]);
  const [newComment, setNewComment] = useState("");
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const fetchData = async () => {
    try {
      const postRes = await api.get(`post/${postId}`);
      setPost(postRes.data);
      const commentsRes = await api.get(`post/${postId}/comment`);
      setComments(commentsRes.data || []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [postId]);

  const handleAddComment = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newComment.trim()) return;
    try {
      await api.post(`user/${userId}/post/${postId}/comment`, {
        text: newComment,
      });
      setNewComment("");
      fetchData();
    } catch (err) {
      console.error(err);
    }
  };

  const handleDeleteComment = async (commentId: number) => {
    if (!window.confirm("Delete this comment?")) return;
    try {
      await api.delete(`user/${userId}/post/${postId}/comment/${commentId}`);
      setComments(comments.filter((c) => c.id !== commentId));
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

  if (!post)
    return (
      <div className="text-center py-20 text-2xl font-black">
        Story not found.
      </div>
    );

  return (
    <div className="min-h-screen bg-white">
      <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-100">
        <div className="container mx-auto px-4 h-16 flex items-center justify-between">
          <button
            onClick={() => navigate(-1)}
            className="flex items-center gap-2 text-slate-500 hover:text-primary-600 font-bold transition-colors"
          >
            <ArrowLeft size={20} />
            <span>Back</span>
          </button>
        </div>
      </nav>

      <main className="container mx-auto px-4 py-16 max-w-4xl">
        <header className="mb-12 text-center">
          <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-primary-50 text-primary-700 text-xs font-black uppercase tracking-widest mb-8">
            <Sparkles size={14} /> Editorial Pick
          </div>
          <h1 className="text-4xl md:text-6xl font-black mb-8 leading-[1.1] tracking-tight text-slate-900">
            {post.title}
          </h1>
          <div className="flex flex-col md:flex-row items-center justify-center gap-6 text-slate-500 font-bold text-sm border-y border-slate-100 py-6">
            <div className="flex items-center gap-3 text-slate-900">
              <UserIcon size={20} />
              <span>{post.author?.name || "Anonymous"}</span>
            </div>
            <div className="flex items-center gap-2">
              <Calendar size={16} />
              {new Date(post.CreatedAt).toLocaleDateString()}
            </div>
          </div>
        </header>

        <article className="prose prose-slate prose-xl max-w-none mb-20">
          <div className="text-lg leading-relaxed text-slate-800 whitespace-pre-wrap">
            {post.content}
          </div>
        </article>

        <section className="bg-slate-50 rounded-3xl p-8 md:p-12">
          <h2 className="text-3xl font-black mb-10">Join the discussion</h2>
          <form onSubmit={handleAddComment} className="mb-12">
            <div className="relative group">
              <textarea
                className="w-full p-6 rounded-2xl border border-slate-200 bg-white shadow-sm focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 focus:outline-none transition-all resize-none text-lg"
                placeholder="What are your thoughts?"
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                rows={4}
              />
              <button
                type="submit"
                className="absolute bottom-4 right-4 bg-primary-600 text-white p-3 rounded-xl shadow-lg shadow-primary-500/30 hover:scale-105 transition-transform"
              >
                <Send size={24} />
              </button>
            </div>
          </form>

          <div className="space-y-6">
            {comments.map((comment) => (
              <div
                key={comment.id}
                className="bg-white p-6 rounded-2xl shadow-sm border border-slate-100 transition-all hover:shadow-md group"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-full bg-slate-100 flex items-center justify-center text-slate-500 font-bold uppercase">
                      {comment.user?.name.charAt(0) || "A"}
                    </div>
                    <div>
                      <div className="text-sm font-black text-slate-900">
                        {comment.user?.name}
                      </div>
                      <div className="text-xs font-bold text-slate-400">
                        {new Date(comment.CreatedAt).toLocaleDateString()}
                      </div>
                    </div>
                  </div>
                  {Number(userId) === comment.user_id && (
                    <button
                      onClick={() => handleDeleteComment(comment.id)}
                      className="p-2 text-slate-300 hover:text-accent-pink transition-colors opacity-0 group-hover:opacity-100"
                    >
                      <Trash2 size={18} />
                    </button>
                  )}
                </div>
                <p className="text-slate-700 leading-relaxed">{comment.text}</p>
              </div>
            ))}
          </div>
        </section>
      </main>
    </div>
  );
};

export default PostDetail;
