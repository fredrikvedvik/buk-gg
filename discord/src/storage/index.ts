import { Storage } from "@google-cloud/storage";

const storage = new Storage({
    projectId: process.env.GCLOUD_PROJECT_ID
});

const bucketName = process.env.GCLOUD_STORAGE_BUCKET

const getBucket = () => {
    if (!bucketName) {
        throw new Error("Missing GCLOUD_STORAGE_BUCKET in env")
    }
    return storage.bucket(bucketName)
}

export const getOrSet = async <T>(key: string, factory: () => Promise<T>) => {
    let item = await get<T>(key)
    if (!item) {
        item = await factory()
        await set(key, item)
    }
    return item;
}

export const set = async <T>(key: string, item: T) => {
    const bucket = getBucket();
    await bucket.file(key).save(JSON.stringify(item))
}

export const get = async <T>(key: string) => {
    try {
        const bucket = getBucket()
        const [content] = await bucket.file(key).download()
        return JSON.parse(content.toString()) as T
    } catch {
        return null
    }
}
