{{SLASH_COMMENTS}}

#ifndef WIDGETS_IMAGE_PROVIDER_H
#define WIDGETS_IMAGE_PROVIDER_H

#include <QDebug>
#include <QMutex>
#include <QQmlExtensionPlugin>
#include <QQuickImageProvider>

class PhotoImageProvider : public QQuickImageProvider {
  // https://doc.qt.io/qt-5/qquickimageprovider.html

public:
  PhotoImageProvider() : QQuickImageProvider(QQuickImageProvider::Image) {}

  void addImage(const QString &id, const QImage &image) {
    if (image.isNull()) {
      qDebug() << "provider: image is null";
      return;
    }

    QMutexLocker locker(&mutex);
    m_images[id] = image;

    qDebug() << "provider: added image id" << id;
  }

  void removeImage(const QString &id) {
    QMutexLocker locker(&mutex);
    m_images.remove(id);

    qDebug() << "provider: removed image id" << id;
  }

  QImage requestImage(const QString &id, QSize *size, const QSize &requestedSize) override {
    QMutexLocker locker(&mutex);
    auto real_id = id.left(id.indexOf("#"));
    QImage image = m_images[real_id];

    if (image.isNull()) {
      qDebug() << "provider: image not found" << real_id;
      image = QImage(QSize(1, 1), QImage::Format_ARGB32);
      image.fill(Qt::transparent);
    } else {
      qDebug() << "provider: get image real_id" << real_id;
      image.scaled(requestedSize);
    }

    if (size)
      *size = requestedSize;

    return image;
  }

private:
  QMutex mutex;
  QHash<QString, QImage> m_images;
};

class PhotoImageManager : public QObject {
  Q_OBJECT;

public:
  explicit PhotoImageManager(QObject *parent = nullptr) : QObject(parent) {
    photoImageProvider = new PhotoImageProvider();
  }

  PhotoImageProvider *photoImageProvider;

signals:
  void imageChanged(const QString &id);

public slots:
  void addImage(const QString &id, const QImage &image) {
    photoImageProvider->addImage(id, image);
    emit imageChanged(id);
  }

  void addImage(const QString &id, const QByteArray &imageByteArray) {
    if (imageByteArray.isEmpty())
      removeImage(id);

    QImage image;
    image.loadFromData(imageByteArray, "JPG");

    if (!image.isNull()) {
      photoImageProvider->addImage(id, image);
      emit imageChanged(id);
    }
  }

  void removeImage(const QString &id) {
    photoImageProvider->removeImage(id);
    emit imageChanged(id);
  }
};

#endif//WIDGETS_IMAGE_PROVIDER_H
