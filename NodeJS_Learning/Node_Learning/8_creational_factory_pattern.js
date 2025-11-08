class Notification {
    send() { }
}

class EmailNotification extends Notification {
    send() {
        console.log('Sending email notification');
    }
}

class SMSNotification extends Notification {
    send() {
        console.log('Sending SMS notification');
    }
}

class NotificationFactory {
    createNotification(type) {
        if (type === 'email') {
            return new EmailNotification();
        } else if (type === 'sms') {
            return new SMSNotification();
        } else {
            throw new Error('Notification type not supported.');
        }
    }
}

const factory = new NotificationFactory();
const emailNotification = factory.createNotification('email');
emailNotification.send();// Output: Sending email notification  const sms
const smsNotification = factory.createNotification('sms');
smsNotification.send(); // Output: Sending SMS notification