'use client';

import { useState, useEffect } from 'react';
import { useUser } from '@clerk/nextjs';
import { Plus } from 'lucide-react';
import { Button } from '../components/ui/button';
import { Card, CardContent } from '../components/ui/card';
import Navbar from '../components/Navbar';
import Footer from '../components/Footer';
import LinkCard from '../components/LinkCard';
import CreateLinkDialog from '../components/CreateLinkDialog';
import LinkDetailDialog from '../components/LinkDetailDialog';
import UpdateUsernameDialog from '../components/UpdateUsernameDialog';
import AppLoader from '@/components/AppLoader';
import Toast from '../components/Toast';
import { userAPI, linkAPI } from '@/lib/api';
import { getSuccessMsg } from '@/lib/success';

// # 'Dashboard' Page Component # 
const Dashboard = () => {
  const [minLoaderDelayDone, setMinLoaderDelayDone] = useState(false);
  const { user, isLoaded } = useUser();
  const [links, setLinks] = useState([]);
  const [selectedLink, setSelectedLink] = useState(null);
  const [userData, setUserData] = useState(null);
  const [isCreateOpen, setIsCreateOpen] = useState(false);
  const [isEditUserOpen, setIsEditUserOpen] = useState(false);
  const [isDetailOpen, setIsDetailOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [toast, setToast] = useState(null);

  useEffect(() => {
    // # Add the minimum 'loader' delay
    const timer = setTimeout(() => {
      setMinLoaderDelayDone(true);
    }, 2000);

    // # Get the 'user' and 'links' data on the component 'mount'
    if (user) {
      getUser();
      getLinks();
    }

    // # Clear the 'timer' on the component 'unmount'
    return () => clearTimeout(timer);

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [user]);

  const showToast = (message, type) => {
    setToast({ message, type });
  };

  // # Get 'User'
  const getUser = async () => {
    try {
      const response = await userAPI.getUser(user.id);
      if (response.success) {
        setUserData(response.payload);
      }
    } catch (err) {
      showToast(err.message, 'error');
    }
  };

  // # Update 'Username'
  const updateUsername = async (username) => {
    try {
      const response = await userAPI.updateUsername(user.id, username);
      if (response.success) {
        const successMsg = getSuccessMsg(response.payload);
        showToast(successMsg, 'success');
        setIsEditUserOpen(false);
        await getUser();
        await getLinks();
      }
    } catch (err) {
      showToast(err.message, 'error');
      setIsEditUserOpen(false);
    }
  };

  // # Create 'Link'
  const createLink = async (data) => {
    try {
      const response = await linkAPI.createLink(user.id, data);
      if (response.success) {
        const successMsg = getSuccessMsg(response.payload);
        showToast(successMsg, 'success');
        setIsCreateOpen(false);
        await getLinks();
      }
    } catch (err) {
      showToast(err.message, 'error');
      setIsCreateOpen(false);
    }
  };

  // # Update 'Link'
  const updateLink = async (linkId, data) => {
    try {
      const response = await linkAPI.updateLink(user.id, linkId, data);
      if (response.success) {
        const successMsg = getSuccessMsg(response.payload);
        showToast(successMsg, 'success');
        setIsDetailOpen(false);
        await getLinks();
      }
    } catch (err) {
      showToast(err.message, 'error');
      setIsDetailOpen(false);
    }
  };

  // # Delete 'Link'
  const deleteLink = async (linkId) => {
    try {
      const response = await linkAPI.deleteLink(linkId);
      if (response.success) {
        const successMsg = getSuccessMsg(response.payload);
        showToast(successMsg, 'success');
        await getLinks();
      }
    } catch (err) {
      showToast(err.message, 'error');
    }
  };

  // # Get 'Link'
  const getLink = async (link) => {
    try {
      const response = await linkAPI.getLink(link._id);
      if (response.success) {
        setSelectedLink(response.payload);
        setIsDetailOpen(true);
      }
    } catch (err) {
      showToast(err.message, 'error');
    }
  };

  // # Get 'Links'
  const getLinks = async () => {
    setLoading(true);
    try {
      const response = await linkAPI.getLinks(user.id);
      if (response.success) {
        setLinks(response.payload);
      }
    } catch (err) {
      showToast(err.message, 'error');
    } finally {
      setLoading(false);
    }
  };

  // # Copy To 'Clipboard'
  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    const message = 'Link copied to clipboard!';
    showToast(message, 'success');
  };

  // # Render the 'app loader' only until the minimum 'loader' delay is 'done' and the 'clerk' is fully 'loaded'
  if (!minLoaderDelayDone || !isLoaded) {
    return <AppLoader />;
  };

  return (
    <>
      <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 flex flex-col transition-colors">
        <Navbar
          username={userData?.username || 'user'}
          onEditUsername={() => setIsEditUserOpen(true)}
        />

        <main className="flex-1 max-w-7xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="flex justify-between items-center mb-8">
            <div>
              <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
                Link Management
              </h1>
              <p className="text-gray-600 dark:text-gray-300 mt-1">
                Create and manage your shortened URLs
              </p>
            </div>

            {/* # 'Create Link' Dialog Component # */}
            <CreateLinkDialog
              open={isCreateOpen}
              onOpenChange={setIsCreateOpen}
              onCreate={createLink}
            />
          </div>

          {loading ? (
            <div className="text-center py-12">
              <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-black dark:border-white mx-auto"></div>
            </div>
          ) : !links || links.length === 0 ? (
            <Card className="text-center py-12 dark:bg-gray-800 dark:border-gray-700">
              <CardContent>
                <svg className="w-16 h-16 text-gray-400 dark:text-gray-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                </svg>
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white mb-2">
                  No links yet
                </h3>
                <p className="text-gray-600 dark:text-gray-400 mb-4">
                  Create your first shortened URL to get started
                </p>
                <Button onClick={() => setIsCreateOpen(true)} className="bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200">
                  <Plus className="w-4 h-4" />
                  Create First Link
                </Button>
              </CardContent>
            </Card>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {links.map((link) => (
                <LinkCard
                  key={link._id}
                  link={link}
                  username={userData?.username || 'user'}
                  onClick={getLink}
                  onDelete={deleteLink}
                  onCopy={copyToClipboard}
                />
              ))}
            </div>
          )}

          {/* # 'Link Detail' Dialog Component # */}
          <LinkDetailDialog
            link={selectedLink}
            username={userData?.username || 'user'}
            open={isDetailOpen}
            onOpenChange={setIsDetailOpen}
            onUpdate={updateLink}
            onDelete={deleteLink}
          />

          {/* # 'Update Username' Dialog Component # */}
          <UpdateUsernameDialog
            open={isEditUserOpen}
            onOpenChange={setIsEditUserOpen}
            currentUsername={userData?.username || 'user'}
            onUpdate={updateUsername}
          />
        </main>

        {/* # 'Footer' Component # */}
        <Footer />

        {/* # 'Toast' Component # */}
        {toast && (
          <Toast
            message={toast.message}
            type={toast.type}
            onClose={() => setToast(null)}
          />
        )}
      </div>
    </>
  );
};

export default Dashboard;
