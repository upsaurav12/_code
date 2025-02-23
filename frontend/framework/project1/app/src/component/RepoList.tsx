'use client'

import { useSession } from 'next-auth/react'
import { useState, useEffect } from 'react'
interface Repository {
    id: number
    name: string
    description: string | null
    html_url: string
    fork: boolean
}

export default function RepoList() {
    const { data: session } = useSession()
    const [repos, setRepos] = useState<Repository[]>([])
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        async function fetchRepos() {
            if (!session?.accessToken) return

            setLoading(true)
            try {
                const response = await fetch('https://api.github.com/user/repos', {
                    headers: {
                        Authorization: `Bearer ${session.accessToken}`,
                        Accept: 'application/vnd.github.v3+json',
                    },
                })

                if (!response.ok) {
                    throw new Error('Failed to fetch repositories')
                }

                const data = await response.json()
                console.log(data)
                setRepos(data)
            } catch (err) {
                setError(err instanceof Error ? err.message : 'An error occurred')
            } finally {
                setLoading(false)
            }
        }

        fetchRepos()
    }, [session])

    if (!session) {
        return <div>Please sign in to view your repositories</div>
    }

    if (loading) return <div>Loading repositories...</div>
    if (error) return <div>Error: {error}</div>

    return (
        <>
            <h1>Hello World</h1>
        </>
    )
}