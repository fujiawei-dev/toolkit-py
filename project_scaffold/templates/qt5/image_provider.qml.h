{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__IMAGE_PROVIDER_H
#define {{APP_NAME_UPPER}}__IMAGE_PROVIDER_H

#include <QMutex>
#include <QMutexLocker>
#include <QQuickImageProvider>

class QmlImageProvider : public QQuickImageProvider {
public:
    QmlImageProvider();

    void updateImage(int index, const QImage &image);
    QImage requestImage(const QString &id, QSize *size, const QSize &requestedSize);

    QHash<int, QImage> images;
    QMutex mutex;
};

class QmlImageManager : public QObject {
    Q_OBJECT;

public:
    explicit QmlImageManager(QObject *parent = nullptr);

    QmlImageProvider *qmlImageProvider;

signals:
    void imageChanged(int index);

public slots:
    void updateImage(int index, const QImage &image);
};

#endif//{{APP_NAME_UPPER}}__IMAGE_PROVIDER_H
