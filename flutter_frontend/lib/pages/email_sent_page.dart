import 'package:flutter/material.dart';

class EmailSentPage extends StatelessWidget {
  const EmailSentPage({
    super.key,
    required this.subject,
    required this.message,
  });

  final String subject;
  final String message;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SizedBox(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                subject,
                style: const TextStyle(fontSize: 32),
              ),
              Text(
                message,
                style: const TextStyle(fontSize: 32),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
